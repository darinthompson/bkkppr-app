import { Box, Button, TextField, Typography } from "@mui/material";
import { useState } from "react";
import './login.css'
import axios from 'axios';
import { useNavigate } from "react-router-dom";
function Auth() {
  const [hasAccount, setHasAccount] = useState(true);
  const [formValues, setFormValues] = useState({
    firstName: "",
    lastName: "",
    username: "",
    email: "",
    password: ""
  });

  const handleChange = (e: any) => {
    const { id, value } = e.target;
    setFormValues((prev) => ({
      ...prev,
      [id]: value,
    }));
  };

  const authenticateUser = async (hasAccount: boolean, userData: unknown) => {
    const endpoint: string = hasAccount ? "/login" : "/signup";

    // Perform API request
    const response = await axios.post(endpoint, userData);

    // Extract data from response
    const { token, user } = response.data;
    console.log("RES: ", response.data);

    // Store token & user data in localStorage
    localStorage.setItem("token", token);
    localStorage.setItem("user", JSON.stringify(user));

    // Navigate to welcome page
    navigate("/welcome");
  }

  const navigate = useNavigate();
  const handleSubmit = () => {
    authenticateUser(hasAccount, {
      username: formValues.username,
      password: formValues.password
    });
  };

  return (
    <Box sx={{padding: 30, display: "inline-flex", flexDirection: "column"}}>
      <Typography color="#d1d5de" variant="h3" sx={{mb:4}}>Welcome To Bkkppr</Typography>
      {!hasAccount && (
        <>
          <TextField
          sx={{ m: 2 }}
          variant="outlined"
          id="firstName"
          label="First Name"
          value={formValues.firstName}
          onChange={handleChange}
          />
          <TextField
          sx={{ m: 2 }}
          variant="outlined"
          id="lastName"
          label="Last Name"
          value={formValues.lastName}
          onChange={handleChange}
        />
        <TextField
          sx={{ m: 2 }}
          variant="outlined"
          id="email"
          label="Email"
          value={formValues.email}
          onChange={handleChange}
        />
      </>
      )}
      <TextField
        sx={{ m: 2 }}
        variant="outlined"
        id="username"
        label="Username"
        value={formValues.username}
        onChange={handleChange}
      />
      <TextField
        type="password"
        sx={{ m: 2 }}
        variant="outlined"
        id="password"
        label="Password"
        value={formValues.password}
        onChange={handleChange}
      />
      <Button variant="contained" onClick={handleSubmit}>
        {hasAccount ? "Log In" : "Sign Up"}
      </Button>
      <Button variant="text" onClick={() => setHasAccount((prev) => !prev)}>
        {hasAccount
          ? "Don't have an account? Sign Up"
          : "Already have an account? Log In"}
      </Button>
    </Box>
  );
}

export default Auth;
