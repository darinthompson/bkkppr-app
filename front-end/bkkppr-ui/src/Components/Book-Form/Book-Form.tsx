import React, { useState } from 'react';
import { TextField, Checkbox, FormControlLabel, Button, Box } from '@mui/material';
import { Book } from '../../models/Book';

interface BookFormValues {
  title: string;
  author: string;
  genre: string;
  publishDate: string;
  isCheckedOut: boolean;
}

const initialFormValues: Book = {
  title: '',
  author: '',
  Genre: '',
  Publish_Date: '',
};

interface BookFormProps {
  addBook: (book: Book) => void;
  closeForm: () => void;
}

const BookForm: React.FC<BookFormProps>  = ({addBook, closeForm}) => {
  const [formValues, setFormValues] = useState<Book>(initialFormValues);

  const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, type, checked } = event.target;
    setFormValues(prevValues => ({
      ...prevValues,
      [name]: type === 'checkbox' ? checked : value,
    }));
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    addBook(formValues);
    console.log(formValues);
    closeForm();
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
        name="Genre"
        label="Genre"
        value={formValues.Genre}
        onChange={handleChange}
      />
      <TextField
        sx={{ m: 2 }}
        variant="outlined"
        id="publishDate"
        name="Publish_Date"
        label="Publish Date"
        value={formValues.Publish_Date}
        onChange={handleChange}
      />
      {/* <FormControlLabel
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
      /> */}
      <Button type="submit" variant="contained" sx={{ m: 2 }}>
        Submit
      </Button>
    </Box>
  );
};

export default BookForm;
