    import { Box, Button, Dialog, DialogActions, DialogContent, DialogTitle, Typography } from "@mui/material";
    import BookCard from "../BookCard";
    import { sampleBooks } from "../../assets/sampleData/BookList";
import { useState } from "react";
import BookForm from "../Book-Form/Book-Form";
    function Home() {

        const [open, setOpen] = useState(false);

        const handleOpen = () => {
          setOpen(true);
        };
      
        const handleClose = () => {
          setOpen(false);
        };

        return (
            <Box sx={{display: "inline-flex", flexDirection: "column"}}>
                <Button variant="contained" onClick={handleOpen} sx={{ m: 2 }}>
                    Add a book
                </Button>
                    <Dialog open={open} onClose={handleClose}>
                        <DialogTitle>Book Form</DialogTitle>
                        <DialogContent>
                        <BookForm />
                        </DialogContent>
                        <DialogActions>
                        <Button onClick={handleClose}>Cancel</Button>
                        <Button onClick={handleClose}>Submit</Button>
                        </DialogActions>
                    </Dialog>
                <Typography color="#d1d5de" variant="h3" sx={{mb:4}}>Welcome to Your Library</Typography>
                {sampleBooks.map((book, index) => (
                    <BookCard
                    key={index}
                    title={book.title}
                    author={book.author}
                    genre={book.genre}
                    publishDate={book.publishDate}
                    isCheckedOut={book.isCheckedOut}
                    />
                ))}
            </Box>
        )
    }

    export default Home;