import {Component, signal} from '@angular/core';
import {Product} from '../../models/products.models';
import {ProductCardComponent} from './product-card/product-card.component';

@Component({
  selector: 'app-products-list',
  standalone: true,
  imports: [
    ProductCardComponent
  ],
  templateUrl: './products-list.component.html',
  styles: ``
})
export class ProductsListComponent {
  products = signal<Product[]>([
    {
      id: 1,
      title: "jktyctytfcjy",
      price: 109.95,
      image: "../hero-image.png",
      stock: 0,
    },
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
    {
      id: 4,
      title: "poiuytrszsxrdfrdf",
      price: 109.95,
      image: "https://image.tmdb.org/t/p/w500",
      stock: 0,
    },
  ])
}
