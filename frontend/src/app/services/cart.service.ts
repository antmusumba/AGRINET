import { Injectable, signal } from '@angular/core';
import { Product } from '../models/products.models';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class CartService {
  private readonly PAYMENT_API_URL =
    'https://payment-gateway.example.com/api/pay';

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
