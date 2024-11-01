# Schema

```mermaid
erDiagram
    LEADERBOARD ||--o{ SCORE : scores
    LEADERBOARD {
        string name
    }
    SCORE {
        int score
        string player
    }

```