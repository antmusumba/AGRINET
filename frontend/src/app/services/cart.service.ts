import { Injectable, signal } from '@angular/core';
import { Product } from '../models/products.models';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../environment/environment.dev';

@Injectable({
  providedIn: 'root',
})
export class CartService {
  private readonly PAYMENT_API_URL = `${environment.apiUrl}/api/paymentgateway`;

  cart = signal<Product[]>([]);

  constructor(private http: HttpClient) {}

  addToCart(product: Product) {
    this.cart.set([...this.cart(), product]);
  }

  removeFromCart(product: Product) {
    this.cart.set(this.cart().filter((item) => item.id !== product.id));
  }

  initiatePayment(phone: string, amount: number): Observable<any> {
    return this.http.post(this.PAYMENT_API_URL, { phone, amount });
  }
}
