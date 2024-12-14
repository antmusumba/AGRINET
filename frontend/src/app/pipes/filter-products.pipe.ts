import { Pipe, PipeTransform } from '@angular/core';
import {Product} from '../models/products.models';

@Pipe({
  standalone: true,
  name: 'filterProducts'
})
export class FilterProductsPipe implements PipeTransform {
  transform(products: Product[], searchQuery: string): Product[] {
    if (!searchQuery) {
      return products; // Return all products if no query
    }
    const lowerCaseQuery = searchQuery.toLowerCase();
    return products.filter((product) =>
      product.title.toLowerCase().includes(lowerCaseQuery)
    );
  }
}
