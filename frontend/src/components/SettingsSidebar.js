import { useLocation, matchPath, useNavigate } from "react-router-dom";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";
import KeyIcon from "@mui/icons-material/Key";
import ReceiptLongIcon from "@mui/icons-material/ReceiptLong";
import Box from "@mui/material/Box";
import Link from "@mui/material/Link";
import InfoOutlinedIcon from "@mui/icons-material/InfoOutlined";
import { Divider } from "@mui/material";

const pages = [
  { name: "Account", href: "/settings", icon: <AccountCircleIcon /> },
  {
    name: "Change Password",
    href: "/settings/changepassword",
    icon: <KeyIcon />,
  },
  {
    name: "Logs",
    href: "/settings/logs",
    icon: <ReceiptLongIcon />,
  },
  {
    divider: true,
  },
  {
    name: "About",
    href: "/settings/about",
    icon: <InfoOutlinedIcon />,
  },
];

const SettingsSidebar = () => {
  const location = useLocation();
  const navigate = useNavigate();

  return (
    <>
      {pages.map((page) =>
        page.divider ? (
          <Divider sx={{ my: 1 }} />
        ) : (
          <Link key={page.href} href={page.href} underline="none">
            <Box
              onClick={(e) => {
                e.preventDefault();
                navigate(page.href);
              }}
              sx={{
                display: "flex",
                flexDirection: "row",
                alignItems: "center",
                justifyContent: "start",
                borderRadius: 1,
                userSelect: "none",
                p: 1,
                my: 0.5,

                background: matchPath(
                  {
                    path: page.href,
                    exact: true,
                    strict: false,
                  },
                  location.pathname
                )
                  ? "#ebebeb"
                  : "unset",

                color: matchPath(
                  {
                    path: page.href,
                    exact: true,
                    strict: false,
                  },
                  location.pathname
                )
                  ? "black"
                  : "gray",

                "&:hover": {
                  background: "#ebf3ff",
                  cursor: "pointer",
                  color: "black",
                },
              }}
            >
              <Box
                sx={{
                  mr: 1,
                  display: "flex",
                  justifyContent: "center",
                  alignItems: "center",
                }}
              >
                {page.icon}
              </Box>
              <Box sx={{ fontSize: 14 }}>{page.name}</Box>
            </Box>
          </Link>
        )
      )}
    </>
  );
};

export default SettingsSidebar;
