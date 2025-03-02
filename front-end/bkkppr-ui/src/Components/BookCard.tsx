// Updated BookCard.tsx
import React from 'react';
import { Card, CardContent, Typography, Box } from '@mui/material';

interface BookCardProps {
  title?: string;
  author?: string;
  genre?: string;
  publish_date?: string;
}

const BookCard: React.FC<BookCardProps> = ({ title, author, genre, publish_date }) => {
  return (
    <Card sx={{ maxWidth: 345, margin: '1rem' }}>
      <CardContent>
        <Typography variant="h5" component="div">
          {title}
        </Typography>
        <Typography sx={{ mb: 1.5 }} color="text.secondary">
          by {author}
        </Typography>
        <Box sx={{ mt: 2 }}>
          <Typography variant="body2">Genre: {genre}</Typography>
          <Typography variant="body2">Published: {publish_date}</Typography>
        </Box>
      </CardContent>
    </Card>
  );
};

export default BookCard;
