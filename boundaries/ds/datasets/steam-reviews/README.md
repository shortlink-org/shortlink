# Steam reviews dataset

### Dataset Structure

The dataset contains the following columns:
- **id**: A unique identifier for each review.
- **game**: The name of the game being reviewed.
- **review**: The text of the Steam review.
- **author_playtime_at_review**: The number of hours the author had played the game at the time of writing the review.
- **voted_up**: Whether the user marked the review/the game as positive (True) or negative (False).
- **votes_up**: The number of upvotes the review received from other users.
- **votes_funny**: The number of "funny" votes the review received from other users.
- **constructive**: A binary label indicating whether the review was constructive (1) or not (0).

### Reference

- [Dataset](https://huggingface.co/datasets/abullard1/steam-reviews-constructiveness-binary-label-annotations-1.5k)
