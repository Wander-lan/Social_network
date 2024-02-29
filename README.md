# A social network app with an API
This is a generic social network application with it's own API and a local database.

# API Routes:

## Login Route
- Endpoint: "/login" (POST)

## Users Manipulation Routes
1. Create User
   - Endpoint: "/users" (POST)
2. Get Users
   - Endpoint: "/users" (GET)
3. Get User by ID
   - Endpoint: "/users/{userId}" (GET)
4. Update User by ID
   - Endpoint: "/users/{userId}" (PUT)
5. Delete User by ID
   - Endpoint: "/users/{userId}" (DELETE)
6. Follow User
   - Endpoint: "/users/{userId}/follow" (POST)
7. Unfollow User
   - Endpoint: "/users/{userId}/unfollow" (POST)
8. Get Followers of User
   - Endpoint: "/users/{userId}/followers" (GET)
9. Get Following Users
   - Endpoint: "/users/{userId}/following" (GET)
10. Update User Password
    - Endpoint: "/users/{userId}/update-password" (POST)

## Publications Manipulation Routes
1. Create Post
   - Endpoint: "/posts" (POST)
2. Get Posts
   - Endpoint: "/posts" (GET)
3. Get Post by ID
   - Endpoint: "/posts/{postId}" (GET)
4. Update Post by ID
   - Endpoint: "/posts/{postId}" (PUT)
5. Delete Post by ID
   - Endpoint: "/posts/{postId}" (DELETE)
6. Get User Posts
   - Endpoint: "/users/{userId}/posts" (GET)
7. Like Post
   - Endpoint: "/posts/{postId}/like" (POST)
8. Dislike Post
   - Endpoint: "/posts/{postId}/dislike" (POST)

