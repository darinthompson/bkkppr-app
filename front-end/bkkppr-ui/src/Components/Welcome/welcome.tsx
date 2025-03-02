    import { Box, Button, Dialog, DialogActions, DialogContent, DialogTitle, Typography } from "@mui/material";
    import BookCard from "../BookCard";
    import { Book } from "../../models/Book";
    import { useEffect, useState } from "react";
    import BookForm from "../Book-Form/Book-Form";
    import axios from "axios";
    
    interface User {
        id: number,
        username: string
    }
    function Home() {

        const [open, setOpen] = useState(false);
        const [user, setUser] = useState<User | null>(() => {
            const storedUser = localStorage.getItem("user");
            return storedUser ? JSON.parse(storedUser) : null; 
        })
        const [books, setBooks] = useState<Book[]>([]);
        const handleOpen = () => {
          setOpen(true);
        };
      
        const handleClose = () => {
          setOpen(false);
        };

        const addBook = (newBook: Book) => {
            newBook.user_id = user?.id;
            setBooks((prevBooks) => [...prevBooks, newBook]);
            axios.post('/books', newBook);
        }

 
        useEffect(() => {
            axios.get<Book[]>('api/users/1/books')
            .then(res => {
                setBooks(res.data);
            });
            
            const storedUser = localStorage.getItem("user");
            if(storedUser) {
                setUser(JSON.parse(storedUser));
            }

        }, [])

        return (
            <Box sx={{display: "inline-flex", flexDirection: "column"}}>
                <Button variant="contained" onClick={handleOpen} sx={{ m: 2 }}>
                    Add a book
                </Button>
                    <Dialog open={open} onClose={handleClose}>
                        <DialogTitle>Book Form</DialogTitle>
                        <DialogContent>
                        <BookForm addBook={addBook} closeForm={() => setOpen(false)}/>
                        </DialogContent>
                        <DialogActions>
                        <Button onClick={handleClose}>Cancel</Button>
                        <Button onClick={handleClose}>Submit</Button>
                        </DialogActions>
                    </Dialog>
                <Typography color="#d1d5de" variant="h3" sx={{mb:4}}>Welcome to Your Library</Typography>

                {books.length > 0 ? (
                    books.map((book, index) => (
                        <BookCard
                            key={index}
                            title={book.title}
                            author={book.author}
                            genre={book.Genre}
                            publish_date={book.Publish_Date}
                        />
                    ))
                ) : (
                    <Typography color="#d1d5de" variant="h5" sx={{mb:4}}>Library is Empty</Typography>
                )}
            </Box>
        )
    }

    export default Home;