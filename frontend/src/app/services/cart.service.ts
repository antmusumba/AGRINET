import { Injectable, signal } from '@angular/core';
import {Product} from '../models/products.models';

@Injectable({
  providedIn: 'root'
})
export class CartService {

  cart = signal<Product[]>([
    {
      id: 2,
      title: "ihyfrdrtdytr",
      price: 109.95,
      image: "",
      stock: 10,
    },
    {
      id: 3,
      title: "iouywertyuio",
      price: 109.95,
      image: "https://image.tmdb.org/t/p/w500",
      stock: 10,
    },
  ]);

  addToCart(product: Product) {
    this.cart.set([...this.cart(), product]);
  }

  removeFromCart(product: Product) {
    this.cart.set(this.cart().filter((item) => item.id !== product.id));
  }

  constructor() { }
}
