import { QueryClient, QueryClientProvider } from "react-query";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import ProjectRoute from "./routes/ProjectRoute";
import React from "react";
import CssBaseline from "@mui/material/CssBaseline";
import Register from "./pages/Register";
import Dashboard from "./pages/Dashboard";
import Login from "./pages/Login";
import AuthorizedRoute from "./routes/AuthorizedRoute";
import Projects from "./pages/Projects";
import ProjectSettings from "./pages/ProjectSettings";
import ProjectFields from "./pages/ProjectFields";
import Account from "./pages/Account";
import ChangePassword from "./pages/ChangePassword";
import Logs from "./pages/Logs";
import About from "./pages/About";
import SettingsRoute from "./routes/SettingsRoute";
import Project from "./pages/Project";
import ProjectAPI from "./pages/ProjectAPI";

const theme = createTheme({
  palette: {
    primary: {
      main: "#4184e9",
    },
    secondary: {
      main: "#edf2ff",
    },
  },
  disableTransition: {
    transition: "none",
  },
});

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 60,
      refetchOnWindowFocus: false,
    },
  },
});

const App = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <ThemeProvider theme={theme}>
          <CssBaseline />

          <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />

            <Route
              path="/projects/"
              element={<AuthorizedRoute element={<Projects />} />}
            />
            <Route
              path="/settings/"
              element={
                <AuthorizedRoute
                  element={<SettingsRoute element={<Account />} />}
                />
              }
            />
            <Route
              path="/settings/changepassword"
              element={
                <AuthorizedRoute
                  element={<SettingsRoute element={<ChangePassword />} />}
                />
              }
            />
            <Route
              path="/settings/logs"
              element={
                <AuthorizedRoute
                  element={<SettingsRoute element={<Logs />} />}
                />
              }
            />
            <Route
              path="/settings/about"
              element={
                <AuthorizedRoute
                  element={<SettingsRoute element={<About />} />}
                />
              }
            />
            <Route
              path="/projects/settings/:projectId"
              element={
                <AuthorizedRoute
                  element={<ProjectRoute element={<ProjectSettings />} />}
                />
              }
            />
            <Route
              path="/projects/:projectId"
              element={
                <AuthorizedRoute
                  element={<ProjectRoute element={<Project />} />}
                />
              }
            />
            <Route
              path="/projects/api/:projectId"
              element={
                <AuthorizedRoute
                  element={<ProjectRoute element={<ProjectAPI />} />}
                />
              }
            />
            <Route
              path="/projects/fields/:projectId"
              element={
                <AuthorizedRoute
                  element={<ProjectRoute element={<ProjectFields />} />}
                />
              }
            />
            <Route
              path="/"
              element={<AuthorizedRoute element={<Dashboard />} />}
            />
          </Routes>
        </ThemeProvider>
      </BrowserRouter>
    </QueryClientProvider>
  );
};

export default App;
