**Project Overview**

DevConnector is a platform for developers to manage their portfolios, share projects, and build a social network with other developers. Developers can showcase their skills, experience, projects, blog posts, and engage with others to receive feedback.

**Key Features:**

**User Authentication and Authorization:**

- Login via OAuth (GitHub integration)
- JWT-based authentication system

**Developer Profile Management:**

- Basic Information (name, profile picture, location, bio)
- Tech Stack (programming languages, frameworks, etc.)
- Experience and Education
- Project Portfolio
- Blog Posts

**Social Networking Features:**

- Developer Following
- Commenting and Liking
- Chatting and Messaging
- Notification System

**Search and Filtering:**

- Tech Stack and Location-based Search
- Keyword Filtering

**Admin Panel:**

- User Management
- Reporting and Blocking

**Technology Stack:**

**Backend:**

- Go (Gin framework) - REST API implementation
- PostgreSQL - Database
- Redis - Caching and real-time features

**Frontend:**

- Next.js - React-based SSR (Server-Side Rendering) and SSG (Static Site Generation)
- TypeScript - Static typing support
- TailwindCSS - Styling and responsive design

**Project Structure:**

**Backend:** `/server`

- `/server/main.go` - Main entry point
- `/server/controllers` - API controllers
- `/server/models` - Data models
- `/server/routes` - Route definitions
- `/server/services` - Business logic

**Frontend:** `/client`

- `/client/pages` - Next.js pages
- `/client/components` - Reusable components
- `/client/styles` - TailwindCSS styles
- `/client/utils` - Utility functions and hooks

