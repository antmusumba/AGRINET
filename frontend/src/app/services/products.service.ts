import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { AuthService } from './auth.service';
import { environment } from '../environment/environment.dev';

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

@Injectable({
  providedIn: 'root',
})
export class ProductsService {
  private readonly API_URL = `${environment.apiUrl}/api/products`;

  constructor(private http: HttpClient, private authService: AuthService) {}

  // Create a new product
  createProduct(product: Partial<Product>): Observable<Product> {
    const userId = this.authService.getUserId();
    const productWithUser = { ...product, userId };
    return this.http.post<Product>(`${this.API_URL}`, productWithUser);
  }

  // Read all products
  getAllProducts(): Observable<Product[]> {
    return this.http.get<Product[]>(`${this.API_URL}`);
  }

  // Read a specific product
  getProductById(productId: string): Observable<Product> {
    return this.http.get<Product>(`${this.API_URL}/${productId}`);
  }

  // Update a product
  updateProduct(
    productId: string,
    updates: Partial<Product>
  ): Observable<Product> {
    return this.http.put<Product>(`${this.API_URL}/${productId}`, updates);
  }

  // Delete a product
  deleteProduct(productId: string): Observable<void> {
    return this.http.delete<void>(`${this.API_URL}/${productId}`);
  }
}
