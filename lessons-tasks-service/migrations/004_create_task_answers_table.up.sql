CREATE TABLE IF NOT EXISTS task_answers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    task_id UUID NOT NULL REFERENCES tasks(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    answer_text TEXT NOT NULL,
    attachment_urls TEXT[] DEFAULT '{}',
    status TEXT NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(task_id, user_id)
);

CREATE INDEX idx_task_answers_task_id ON task_answers(task_id);
CREATE INDEX idx_task_answers_user_id ON task_answers(user_id);