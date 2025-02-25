import React, { useState } from 'react';
import { TextField, Checkbox, FormControlLabel, Button, Box } from '@mui/material';
import axios from 'axios';

interface BookFormValues {
  title: string;
  author: string;
  genre: string;
  publishDate: string;
  isCheckedOut: boolean;
}

const initialFormValues: BookFormValues = {
  title: '',
  author: '',
  genre: '',
  publishDate: '',
  isCheckedOut: false,
};

const BookForm: React.FC = () => {
  const [formValues, setFormValues] = useState<BookFormValues>(initialFormValues);

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, type, checked } = event.target;
    setFormValues(prevValues => ({
      ...prevValues,
      [name]: type === 'checkbox' ? checked : value,
    }));
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    axios.post('/books', formValues);
    console.log(formValues);
  };

  return (
    <Box
      component="form"
      onSubmit={handleSubmit}
      noValidate
      sx={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        p: 2,
      }}
    >
      <TextField
        sx={{ m: 2 }}
        variant="outlined"
        id="title"
        name="title"
        label="Title"
        value={formValues.title}
        onChange={handleChange}
      />
      <TextField
        sx={{ m: 2 }}
        variant="outlined"
        id="author"
        name="author"
        label="Author"
        value={formValues.author}
        onChange={handleChange}
      />
      <TextField
        sx={{ m: 2 }}
        variant="outlined"
        id="genre"
        name="genre"
        label="Genre"
        value={formValues.genre}
        onChange={handleChange}
      />
      <TextField
        sx={{ m: 2 }}
        variant="outlined"
        id="publishDate"
        name="publishDate"
        label="Publish Date"
        value={formValues.publishDate}
        onChange={handleChange}
      />
      <FormControlLabel
        sx={{ m: 2 }}
        control={
          <Checkbox
            id="isCheckedOut"
            name="isCheckedOut"
            checked={formValues.isCheckedOut}
            onChange={handleChange}
          />
        }
        label="Checked Out"
      />
      <Button type="submit" variant="contained" sx={{ m: 2 }}>
        Submit
      </Button>
    </Box>
  );
};

export default BookForm;
