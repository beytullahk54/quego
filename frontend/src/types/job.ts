// src/types/user.ts
export interface Job {
    id: number;
    url: string;
    method: string;
    execute_at: string;
    status: string;
    retry_count: number;
    max_retries: number;
    token_id: string;
    created_at: string;
  }