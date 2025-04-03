CREATE TABLE IF NOT EXISTS user_tasks (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    task_id UUID NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    completed BOOLEAN DEFAULT false,
    completed_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY (user_id, task_id)
);

CREATE INDEX idx_user_tasks_user_id ON user_tasks(user_id);
CREATE INDEX idx_user_tasks_task_id ON user_tasks(task_id);
