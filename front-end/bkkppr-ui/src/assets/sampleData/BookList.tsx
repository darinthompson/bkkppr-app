// sampleBooks.ts
export interface Book {
    title: string;
    author: string;
    genre: string;
    publishDate: string;
    isCheckedOut: boolean;
  }
  
  export const sampleBooks: Book[] = [
    {
      title: "The Great Gatsby",
      author: "F. Scott Fitzgerald",
      genre: "Classic Fiction",
      publishDate: "1925",
      isCheckedOut: false,
    },
    {
      title: "1984",
      author: "George Orwell",
      genre: "Dystopian",
      publishDate: "1949",
      isCheckedOut: true,
    },
    {
      title: "To Kill a Mockingbird",
      author: "Harper Lee",
      genre: "Fiction",
      publishDate: "1960",
      isCheckedOut: false,
    },
    {
      title: "Pride and Prejudice",
      author: "Jane Austen",
      genre: "Romance",
      publishDate: "1813",
      isCheckedOut: false,
    },
    {
      title: "Moby Dick",
      author: "Herman Melville",
      genre: "Adventure",
      publishDate: "1851",
      isCheckedOut: true,
    },
    {
      title: "The Catcher in the Rye",
      author: "J.D. Salinger",
      genre: "Fiction",
      publishDate: "1951",
      isCheckedOut: false,
    },
    {
      title: "The Hobbit",
      author: "J.R.R. Tolkien",
      genre: "Fantasy",
      publishDate: "1937",
      isCheckedOut: false,
    },
    {
      title: "Fahrenheit 451",
      author: "Ray Bradbury",
      genre: "Dystopian",
      publishDate: "1953",
      isCheckedOut: true,
    },
    {
      title: "War and Peace",
      author: "Leo Tolstoy",
      genre: "Historical Fiction",
      publishDate: "1869",
      isCheckedOut: false,
    },
    {
      title: "The Odyssey",
      author: "Homer",
      genre: "Epic",
      publishDate: "8th century BC",
      isCheckedOut: false,
    },
    {
      title: "Crime and Punishment",
      author: "Fyodor Dostoevsky",
      genre: "Psychological Fiction",
      publishDate: "1866",
      isCheckedOut: true,
    },
    {
      title: "Brave New World",
      author: "Aldous Huxley",
      genre: "Dystopian",
      publishDate: "1932",
      isCheckedOut: false,
    },
    {
      title: "The Lord of the Rings",
      author: "J.R.R. Tolkien",
      genre: "Fantasy",
      publishDate: "1954",
      isCheckedOut: true,
    },
  ];
  