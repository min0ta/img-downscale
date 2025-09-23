CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    original_url TEXT NOT NULL,
    current_status VARCHAR(10) NOT NULL,
    result_url TEXT NOT NULL DEFAULT '',
    error_msg TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_task_status ON tasks(current_status);
