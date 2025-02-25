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

  const navigate = useNavigate();
  const handleSubmit = () => {
    console.log(formValues);
    if(hasAccount) {
      axios.post('/login', {
        userName: formValues.username,
        password: formValues.password
      }).then((res) => {
        console.log('RES: ', res);
        navigate('/welcome');
      })
    } else {
      axios.post('/signup', formValues).then((res) => {
        console.log('RES: ', res.data);
        navigate('/welcome')
      })
    }

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
