Implement Favorites page for the Book Management project.

1. Create favorite_books table with columns [user_id, book_id, created_at]

2. Get user_id from JWT

3. implement endpoint :
   - GET /books/favorites - Returns a paginated list of the user’s favorite books
   - PUT /books/{bookId}/favorites - Add a book to favorites
   - DELETE /books/{bookId}/favorites - Remove a book from favorites
