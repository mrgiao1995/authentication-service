insert into auth.users(
  user_name, 
  password, 
  role, 
  is_active)
values (
  'admin',
  '123456',
  'admin',
  true
),(
  'user',
  '123456',
  'user',
  true
)