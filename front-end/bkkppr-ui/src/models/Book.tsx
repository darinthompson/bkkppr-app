export interface Book {
    id?: number;  // Optional because Gorm auto-generates it
    title: string;
    author?: string;
    isbn?: string;
    user_id?: number;
    // owner: User;
    Genre: string;
    Publish_Date: string;
  }