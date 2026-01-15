# 🕵️‍♂️ Impostor

A modern, real-time social deduction game inspired by Spyfall and The Chameleon. Built with **Go** (Fiber) and **SvelteKit/Astro**.

## ✨ Features

-   **Online Multiplayer**: Create lobbies and play with friends in real-time.
-   **Offline "Pass & Play"**: Play locally on a single device.
    -   Configurable Difficulty: **Normal** (Impostor knows nothing) or **Easy** (Impostor gets a hint/trap word).
    -   Random Starting Player: The game decides who starts the interrogation.
-   **Bilingual Support**: Full English 🇺🇸 and Spanish 🇪🇸 translations.
-   **Responsive UI**: Sleek, mobile-first design with glassmorphism aesthetics.

## 🛠️ Tech Stack

-   **Backend**: Go (Golang) using [Fiber](https://gofiber.io/) web framework.
-   **Frontend**: [Astro](https://astro.build/) + [Svelte](https://svelte.dev/) + [TailwindCSS](https://tailwindcss.com/).
-   **State Management**: Svelte Stores + Cookies (for offline persistence).

## 🚀 Getting Started

### Prerequisites

-   [Go](https://go.dev/) (v1.18+)
-   [Node.js](https://nodejs.org/) (v16+) & npm

### Installation

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/Danimarqz/impostor.git
    cd impostor
    ```

2.  **Start the Backend Server**:
    ```bash
    go run ./cmd/server/main.go
    ```
    *Server runs on `http://localhost:8080`*

3.  **Start the Frontend**:
    ```bash
    cd web
    npm install
    npm run dev
    ```
    *Frontend runs on `http://localhost:4321`*

## 🎮 How to Play

1.  **Select a Mode**: Choose **"Join Mission"** for online play or **"Play Offline"** for local play.
2.  **Setup**:
    -   **Online**: Enter your name and Join/Create a lobby. Share the code with friends.
    -   **Offline**: Select category, player count, and difficulty.
3.  **Receive Roles**:
    -   **Civilians**: Receive a secret word (e.g., "Beach").
    -   **Impostor**: Receives "Impostor" (or a trap word in Easy mode).
4.  **Interrogate**: Players take turns asking questions related to the secret word to find out who doesn't know it.
    -   *Tip*: Questions shouldn't be too obvious!
5.  **Vote & Eliminate**: After a set time or when ready, vote to eliminate the suspected Impostor.

## 📜 License

MIT
