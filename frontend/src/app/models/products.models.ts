export interface Product {
  id: string;
  userId: string;
  title: string;
  image: string;
  description: string;
  price: number;
  createdAt: string;
  updatedAt: string;
  stock: number;
}

export interface User {
  id: string;
  email: string;
  firstName: string;
  lastName: string;
  phone: string;
}

export interface AuthResponse {
  user: User;
  token: string;
}

export interface RegisterRequest {
  firstName: string;
  lastName: string;
  email: string;
  phone: string;
  password: string;
}

export interface ProductResponse {
  status: string;
  message: string;
  data: Product[];
}
