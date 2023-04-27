import { useLocation, matchPath, useNavigate } from "react-router-dom";
import { useState, cloneElement } from "react";
import { useLogout } from "../hooks/useLogout";
import AppBar from "@mui/material/AppBar";
import Link from "@mui/material/Link";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import Menu from "@mui/material/Menu";
import MenuIcon from "@mui/icons-material/Menu";
import MenuItem from "@mui/material/MenuItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import Container from "@mui/material/Container";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import Tooltip from "@mui/material/Tooltip";
import KeyboardArrowDownIcon from "@mui/icons-material/KeyboardArrowDown";
import Settings from "@mui/icons-material/Settings";
import Logout from "@mui/icons-material/Logout";
import Logo from "./Logo";
import ProjectListMenu from "./ProjectListMenu";

const pages = [
  { name: "Dashboard", href: "/", match: "/" },
  {
    name: "Projects",
    href: "/projects",
    match: "/projects/*",
    subMenu: <ProjectListMenu />,
  },
  { name: "Organizations", href: "/organizations", match: "/organizations/*" },
];

const NavBar = () => {
  const location = useLocation();
  const navigate = useNavigate();

  const [anchorElNav, setAnchorElNav] = useState(null);
  const [anchorElUser, setAnchorElUser] = useState(null);
  const [anchorElSubMenu, setAnchorElSubMenu] = useState(null);

  const [logout, logoutContext] = useLogout();

  const handleOpenNavMenu = (event) => {
    setAnchorElNav(event.currentTarget);
  };

  const handleOpenUserMenu = (event) => {
    setAnchorElUser(event.currentTarget);
  };

  const handleOpenSubMenu = (event) => {
    setAnchorElSubMenu(event.currentTarget);
  };

  const handleCloseNavMenu = () => {
    setAnchorElNav(null);
  };

  const handleCloseUserMenu = () => {
    setAnchorElUser(null);
  };

  const handleCloseSubMenu = () => {
    setAnchorElSubMenu(null);
  };

  const handleOpenUserSettings = (event) => {
    event.preventDefault();

    navigate("/settings");
    handleCloseUserMenu();
  };

  const handleLogout = () => {
    logout().then(() => {
      navigate("/login");
    });
  };

  return (
    <AppBar
      position="static"
      sx={{
        background: "white",
        boxShadow: "0px 0px 4px 0px #00000033",
        zIndex: 1,
      }}
    >
      <Container>
        <Toolbar disableGutters sx={{ height: 0 }}>
          <Box sx={{ flexGrow: 1, display: { xs: "flex", md: "none" } }}>
            <IconButton size="large" onClick={handleOpenNavMenu}>
              <MenuIcon />
            </IconButton>
            <Menu
              id="menu-appbar"
              anchorEl={anchorElNav}
              anchorOrigin={{
                vertical: "bottom",
                horizontal: "left",
              }}
              keepMounted
              transformOrigin={{
                vertical: "top",
                horizontal: "left",
              }}
              open={Boolean(anchorElNav)}
              onClose={handleCloseNavMenu}
              sx={{
                display: { xs: "block", md: "none" },
              }}
            >
              {pages.map((page) => (
                <Link
                  href={page.href}
                  underline="none"
                  sx={{ color: "unset" }}
                  key={page.name}
                >
                  <MenuItem
                    onClick={(e) => {
                      e.preventDefault();
                      navigate(page.href);
                      handleCloseNavMenu();
                    }}
                  >
                    <Typography textAlign="center">{page.name}</Typography>
                  </MenuItem>
                </Link>
              ))}
            </Menu>
          </Box>

          <Box sx={{ display: "flex" }}>
            <Link href={"/"} underline="none">
              <Button
                onClick={(e) => {
                  e.preventDefault();
                  navigate("/");
                }}
                sx={{
                  color: "black",
                  mr: 1,
                  textTransform: "lowercase",
                  "&:hover": {
                    background: "#00000033",
                  },
                }}
              >
                <Logo />
              </Button>
            </Link>
          </Box>

          <Box
            sx={{
              justifyContent: "left",
              display: { xs: "none", md: "flex" },
              color: "black",
            }}
          >
            {pages.map((page) => (
              <Box
                key={page.name}
                sx={{
                  boxShadow: matchPath(
                    {
                      path: page.match,
                      exact: true,
                      strict: false,
                    },
                    location.pathname
                  )
                    ? "inset 0px -4px 0px 0px #4184e9;"
                    : "",
                }}
              >
                <Link href={page.href} underline="none">
                  <Button
                    onClick={(e) => {
                      e.preventDefault();

                      page.subMenu ? handleOpenSubMenu(e) : navigate(page.href);
                    }}
                    endIcon={
                      page.subMenu && (
                        <KeyboardArrowDownIcon
                          sx={{ ml: -0.5, color: "gray" }}
                        />
                      )
                    }
                    sx={{
                      my: 2,
                      textTransform: "none",
                      color: "black",

                      "&:hover": {
                        background: "#00000022",
                      },
                    }}
                  >
                    {page.name}
                  </Button>
                </Link>

                {page.subMenu && (
                  <Menu
                    sx={{ mt: "40px" }}
                    anchorEl={anchorElSubMenu}
                    anchorOrigin={{
                      vertical: "top",
                      horizontal: "left",
                    }}
                    transformOrigin={{
                      vertical: "top",
                      horizontal: "left",
                    }}
                    open={Boolean(anchorElSubMenu)}
                    onClose={handleCloseSubMenu}
                  >
                    {cloneElement(page.subMenu, {
                      closeSubMenu: handleCloseSubMenu,
                    })}
                  </Menu>
                )}
              </Box>
            ))}
          </Box>

          <Box
            sx={{
              display: "flex",
              flexGrow: 1,
              justifyContent: "right",
            }}
          >
            <Tooltip title="Open settings">
              <IconButton onClick={handleOpenUserMenu} sx={{ p: 0 }}>
                <Avatar sx={{ backgroundColor: "#e3e3e3", color: "#585858" }} />
              </IconButton>
            </Tooltip>
            <Menu
              sx={{ mt: "45px" }}
              anchorEl={anchorElUser}
              anchorOrigin={{
                vertical: "top",
                horizontal: "right",
              }}
              keepMounted
              transformOrigin={{
                vertical: "top",
                horizontal: "right",
              }}
              open={Boolean(anchorElUser)}
              onClose={handleCloseUserMenu}
            >
              <Link href="/settings" underline="none" sx={{ color: "unset" }}>
                <MenuItem onClick={handleOpenUserSettings}>
                  <ListItemIcon>
                    <Settings fontSize="small" />
                  </ListItemIcon>
                  Settings
                </MenuItem>
              </Link>
              <MenuItem
                onClick={handleLogout}
                disabled={logoutContext.isLoading}
              >
                <ListItemIcon>
                  <Logout fontSize="small" />
                </ListItemIcon>
                Sign out
              </MenuItem>
            </Menu>
          </Box>
        </Toolbar>
      </Container>
    </AppBar>
  );
};

export default NavBar;
