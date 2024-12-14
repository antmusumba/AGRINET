import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { map, Observable } from 'rxjs';
import { environment } from '../environment/environment.dev';
import { AuthService } from './auth.service';
import { Product } from '../models/products.models';

interface ProductResponse {
  status: string;
  message: string;
  data: Product[];
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
    return this.http.get<ProductResponse>(`${this.API_URL}`).pipe(
      map((response) => response.data) // Extract the 'data' array
    );
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
