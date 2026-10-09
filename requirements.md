# Game Application

---

# Use Cases

---
## User Use Cases

### Register
user can register to application by phone number

### Login
user can log in to the application by phone number and password

## Game Use Cases
### Each game have a given number of questions
### The difficulty level of questions are "Easy", "Medium" and "Hard"
### Game winner is determined by number of correct answers that each user answered
### Each game belongs to a specific category: sport, history, etc

---

# Entity

## User
- ID
- Phone number
- Avatar
- Name

## Game
- ID
- Category
- Question List
- Players

## Question
- ID
- Question
- Answer List
- Correct Answer
- Difficulty
- Category

## Category
- ID
- Name

