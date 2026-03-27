<img width="1321" height="205" alt="Screenshot 2026-03-27 at 21 34 59" src="https://github.com/user-attachments/assets/e02fa8df-0d20-4526-af09-354d866f4f6a" />
<img width="1386" height="205" alt="Screenshot 2026-03-27 at 21 35 15" src="https://github.com/user-attachments/assets/943a3fac-625a-4e22-ad66-0fda8bee61f8" />
curl -X POST http://localhost:8080/authors -H "Content-Type: application/json" -d '{"name": "Stephen King"}'
{"id":1,"name":"Stephen King"}%          
curl -X POST http://localhost:8080/categories -H "Content-Type: application/json" -d '{"name": "Fiction"}'
{"id":1,"name":"Fiction"}%     
curl -X POST http://localhost:8080/books -H "Content-Type: application/json" -d '{"title": "The Shining", "author_id": 1, "category_id": 1, "price": 19.99}'
{"id":1,"title":"The Shining","author_id":1,"category_id":1,"price":19.99}%                                                                                            
curl -X POST http://localhost:8080/books -H "Content-Type: application/json" -d '{"title": "Bad Book", "author_id": 1, "category_id": 1, "price": -5.00}'
i have some issues with postman because of this i check through the terminal
