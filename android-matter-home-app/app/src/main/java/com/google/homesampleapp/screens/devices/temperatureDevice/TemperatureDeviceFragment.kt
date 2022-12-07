/*
 * Copyright 2022 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package com.google.homesampleapp.screens.devices.temperatureDevice

import android.content.res.Resources
import android.graphics.drawable.Drawable
import android.os.Bundle
import android.text.Html
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.activity.result.ActivityResultLauncher
import androidx.activity.result.IntentSenderRequest
import androidx.databinding.DataBindingUtil
import androidx.fragment.app.Fragment
import androidx.fragment.app.activityViewModels
import androidx.fragment.app.viewModels
import androidx.navigation.findNavController
import androidx.navigation.fragment.findNavController
import com.google.android.material.dialog.MaterialAlertDialogBuilder
import com.google.homesampleapp.ALLOW_DEVICE_SHARING_ON_DUMMY_DEVICE
import com.google.homesampleapp.DeviceState
import com.google.homesampleapp.ON_OFF_SWITCH_DISABLED_WHEN_DEVICE_OFFLINE
import com.google.homesampleapp.OPEN_COMMISSIONING_WINDOW_API
import com.google.homesampleapp.OpenCommissioningWindowApi
import com.google.homesampleapp.PERIODIC_UPDATE_INTERVAL_DEVICE_SCREEN_SECONDS
import com.google.homesampleapp.R
import com.google.homesampleapp.TaskStatus.InProgress
import com.google.homesampleapp.data.DevicesStateRepository
import com.google.homesampleapp.databinding.FragmentDeviceBinding
import com.google.homesampleapp.displayString
import com.google.homesampleapp.formatTimestamp
import com.google.homesampleapp.isDummyDevice
import com.google.homesampleapp.lifeCycleEvent
import com.google.homesampleapp.screens.shared.SelectedDeviceViewModel
import com.google.homesampleapp.stateDisplayString
import dagger.hilt.android.AndroidEntryPoint
import javax.inject.Inject
import timber.log.Timber

/**
 * The Device Fragment shows all the information about the device that was selected in the Home
 * screen and supports the following actions:
 * ```
 * - toggle the on/off state of the device
 * - share the device with another Matter commissioner app
 * - inspect the device (get all info we can from the clusters supported by the device)
 * ```
 * When the Fragment is viewable, a periodic ping is sent to the device to get its latest
 * information. Main use case is to update the device's online status dynamically.
 */
@AndroidEntryPoint
class TemperatureDeviceFragment : Fragment() {

  @Inject internal lateinit var devicesStateRepository: DevicesStateRepository

  // Fragment binding.
  private lateinit var binding: FragmentDeviceBinding

  // The ViewModel for the currently selected device.
  private val selectedDeviceViewModel: SelectedDeviceViewModel by activityViewModels()

  // The fragment's ViewModel.
  private val viewModel: DeviceViewModel by viewModels()

  // -----------------------------------------------------------------------------------------------
  // Lifecycle functions

  override fun onCreateView(
      inflater: LayoutInflater,
      container: ViewGroup?,
      savedInstanceState: Bundle?
  ): View {
    super.onCreateView(inflater, container, savedInstanceState)

    Timber.d(lifeCycleEvent("onCreateView()"))

    // Setup the binding with the fragment.
    binding = DataBindingUtil.inflate(inflater, R.layout.fragment_device_temperature, container, false)

    // Setup UI elements and livedata observers.
    setupUiElements()
    setupObservers()

    return binding.root
  }

  override fun onResume() {
    super.onResume()
    Timber.d("onResume(): Starting periodic ping on device")
    if (PERIODIC_UPDATE_INTERVAL_DEVICE_SCREEN_SECONDS != -1) {
      viewModel.startDevicePeriodicPing(selectedDeviceViewModel.selectedDeviceLiveData.value!!)
    }
  }

  override fun onPause() {
    super.onPause()
    Timber.d("onPause(): Stopping periodic ping on device")
    viewModel.stopDevicePeriodicPing()
  }

  // -----------------------------------------------------------------------------------------------
  // Setup UI elements

  private fun setupUiElements() {
    // Navigate back
    binding.topAppBar.setOnClickListener {
      Timber.d("topAppBar.setOnClickListener")
      findNavController().popBackStack()
    }

    // Info / Inspect device menu button
    val description = getString(R.string.inspect_description)
    val descriptionHtml = Html.fromHtml(description, Html.FROM_HTML_MODE_LEGACY)
    binding.topAppBar.setOnMenuItemClickListener {
      MaterialAlertDialogBuilder(requireContext())
          .setTitle(
              getString(
                  R.string.inspect_device_name,
                  selectedDeviceViewModel.selectedDeviceLiveData.value?.device?.name))
          .setMessage(descriptionHtml)
          .setPositiveButton(resources.getString(R.string.ok)) { _, _ ->
            viewModel.inspectDescriptorCluster(
                selectedDeviceViewModel.selectedDeviceLiveData.value!!)
          }
          .show()
      true
    }

    // Remove Device
    binding.removeButton.setOnClickListener {
      val deviceId = selectedDeviceViewModel.selectedDeviceIdLiveData.value
      MaterialAlertDialogBuilder(requireContext())
          .setTitle("Remove this device?")
          .setMessage(
              "This device will be removed and unlinked from this sample app. Other services and connection-types may still have access.")
          .setNegativeButton(resources.getString(R.string.cancel)) { _, _ ->
            // nothing to do
          }
          .setPositiveButton(resources.getString(R.string.yes_remove_it)) { _, _ ->
            // TODO: the log message below never shows up, don't know why.
            selectedDeviceViewModel.resetSelectedDevice()
            val bundle = Bundle()
            bundle.putString("snackbarMsg", "Removing device [${deviceId}]")
            // This must be done before calling viewModel.removeDevice() otherwise the bundle won't
            // make it in HomeFragment.onViewCreated().
            // TODO: to be investigated
            requireView()
                .findNavController()
                .navigate(R.id.action_deviceFragment_to_homeFragment, bundle)
            viewModel.removeDevice(deviceId!!)
          }
          .show()
    }
  }

  // -----------------------------------------------------------------------------------------------
  // Setup Observers

  private fun setupObservers() {
    // Generic status about actions processed in this screen.
    devicesStateRepository.lastUpdatedDeviceState.observe(viewLifecycleOwner) {
      Timber.d(
          "devicesStateRepository.lastUpdatedDeviceState.observe: [${devicesStateRepository.lastUpdatedDeviceState.value}]")
      updateDeviceInfo(devicesStateRepository.lastUpdatedDeviceState.value)
    }


    // Observer on the currently selected device
    selectedDeviceViewModel.selectedDeviceIdLiveData.observe(viewLifecycleOwner) { deviceId ->
      Timber.d(
          "selectedDeviceViewModel.selectedDeviceIdLiveData.observe is called with deviceId [${deviceId}]")
      updateDeviceInfo(null)
    }
  }

  // -----------------------------------------------------------------------------------------------
  // UI update functions

  private fun updateDeviceInfo(deviceState: DeviceState?) {
    if (selectedDeviceViewModel.selectedDeviceIdLiveData.value == -1L) {
      // Device was just removed, nothing to do. We'll move to HomeFragment.
      return
    }
    val deviceUiModel = selectedDeviceViewModel.selectedDeviceLiveData.value

    // Device state
    deviceUiModel?.let {
      val isOnline =
          when (deviceState) {
            null -> deviceUiModel.isOnline
            else -> deviceState.online
          }
      val isOn =
          when (deviceState) {
            null -> deviceUiModel.isOn
            else -> deviceState.on
          }

      binding.topAppBar.title = deviceUiModel.device.name
      binding.onOffTextView.text = stateDisplayString(isOnline, isOn)
      binding.techInfoDetailsTextView.text =
          getString(
              R.string.share_device_info,
              formatTimestamp(deviceUiModel.device.dateCommissioned!!, null),
              deviceUiModel.device.deviceId.toString(),
              deviceUiModel.device.vendorId,
              deviceUiModel.device.productId,
              deviceUiModel.device.deviceType.displayString())
    }
  }

}
