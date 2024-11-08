// sampleData.js
export const users = [];
export const users1 = [
  {
    id: 1,
    username: "johndoe",
    email: "johndoe@example.com",
    hashed_master_password: "hashed_example_1", // Just a placeholder
    salt: "salt_example_1",
    created_at: "2024-11-03T10:00:00Z",
  },
  {
    id: 2,
    username: "janedoe",
    email: "janedoe@example.com",
    hashed_master_password: "hashed_example_2",
    salt: "salt_example_2",
    created_at: "2024-11-03T11:00:00Z",
  },
];

export const passwords = [];
export const passwords1 = [
  {
    id: 1,
    website: "example.com",
    description: "Personal email account",
    email: "john@example.com",
    username: "john.doe",
    password: "encrypted_password_1", // Placeholder for encrypted password
    access_count: 5,
    created_at: "2024-11-01T10:00:00Z",
    updated_at: "2024-11-02T10:00:00Z",
    user_id: 1,
  },
  {
    id: 2,
    website: "work.com",
    description: "Work account login",
    email: "jane@work.com",
    username: "jane.doe",
    password: "encrypted_password_2",
    access_count: 2,
    created_at: "2024-11-02T12:00:00Z",
    updated_at: "2024-11-03T09:00:00Z",
    user_id: 2,
  },
  {
    id: 3,
    website: "socialmedia.com",
    description: "Social media account",
    email: "john@social.com",
    username: "john123",
    password: "encrypted_password_3",
    access_count: 10,
    created_at: "2024-10-29T08:30:00Z",
    updated_at: "2024-11-03T07:45:00Z",
    user_id: 1,
  },
  {
    id: 4,
    website: "socialmedia.com",
    description: "Social media account",
    email: "john@social.com",
    username: "john123",
    password: "encrypted_password_3",
    access_count: 10,
    created_at: "2024-10-29T08:30:00Z",
    updated_at: "2024-11-03T07:45:00Z",
    user_id: 1,
  },
  {
    id: 1,
    website: "example.com",
    description: "Personal email account",
    email: "john@example.com",
    username: "john.doe",
    password: "encrypted_password_1", // Placeholder for encrypted password
    access_count: 5,
    created_at: "2024-11-01T10:00:00Z",
    updated_at: "2024-11-02T10:00:00Z",
    user_id: 1,
  },
  {
    id: 2,
    website: "work.com",
    description: "Work account login",
    email: "jane@work.com",
    username: "jane.doe",
    password: "encrypted_password_2",
    access_count: 2,
    created_at: "2024-11-02T12:00:00Z",
    updated_at: "2024-11-03T09:00:00Z",
    user_id: 2,
  },
  {
    id: 3,
    website: "socialmedia.com",
    description: "Social media account",
    email: "john@social.com",
    username: "john123",
    password: "encrypted_password_3",
    access_count: 10,
    created_at: "2024-10-29T08:30:00Z",
    updated_at: "2024-11-03T07:45:00Z",
    user_id: 1,
  },
  {
    id: 4,
    website: "socialmedia.com",
    description: "Social media account",
    email: "john@social.com",
    username: "john123",
    password: "encrypted_password_3",
    access_count: 10,
    created_at: "2024-10-29T08:30:00Z",
    updated_at: "2024-11-03T07:45:00Z",
    user_id: 1,
  },
  {
    id: 5,
    website: "socialmedia.com",
    description: "Social media account",
    email: "john@social.com",
    username: "john123",
    password: "encrypted_password_3",
    access_count: 10,
    created_at: "2024-10-29T08:30:00Z",
    updated_at: "2024-11-03T07:45:00Z",
    user_id: 1,
  },
  {
    id: 6,
    website: "socialmedia.com",
    description: "Social media account",
    email: "john@social.com",
    username: "john123",
    password: "encrypted_password_3",
    access_count: 10,
    created_at: "2024-10-29T08:30:00Z",
    updated_at: "2024-11-03T07:45:00Z",
    user_id: 1,
  },
  {
    id: 1,
    website: "example.com",
    description: "Personal email account",
    email: "john@example.com",
    username: "john.doe",
    password: "encrypted_password_1", // Placeholder for encrypted password
    access_count: 5,
    created_at: "2024-11-01T10:00:00Z",
    updated_at: "2024-11-02T10:00:00Z",
    user_id: 1,
  },
  {
    id: 2,
    website: "work.com",
    description: "Work account login",
    email: "jane@work.com",
    username: "jane.doe",
    password: "encrypted_password_2",
    access_count: 2,
    created_at: "2024-11-02T12:00:00Z",
    updated_at: "2024-11-03T09:00:00Z",
    user_id: 2,
  },
  {
    id: 3,
    website: "socialmedia.com",
    description: "Social media account",
    email: "john@social.com",
    username: "john123",
    password: "encrypted_password_3",
    access_count: 10,
    created_at: "2024-10-29T08:30:00Z",
    updated_at: "2024-11-03T07:45:00Z",
    user_id: 1,
  },
  {
    id: 4,
    website: "socialmedia.com",
    description: "Social media account",
    email: "john@social.com",
    username: "john123",
    password: "encrypted_password_3",
    access_count: 10,
    created_at: "2024-10-29T08:30:00Z",
    updated_at: "2024-11-03T07:45:00Z",
    user_id: 1,
  },
  {
    id: 5,
    website: "socialmedia.com",
    description: "Social media account",
    email: "john@social.com",
    username: "john123",
    password: "encrypted_password_3",
    access_count: 10,
    created_at: "2024-10-29T08:30:00Z",
    updated_at: "2024-11-03T07:45:00Z",
    user_id: 1,
  },
  {
    id: 6,
    website: "socialmedia.com",
    description: "Social media account",
    email: "john@social.com",
    username: "john123",
    password: "encrypted_password_3",
    access_count: 10,
    created_at: "2024-10-29T08:30:00Z",
    updated_at: "2024-11-03T07:45:00Z",
    user_id: 1,
  },
];
