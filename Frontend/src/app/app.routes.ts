import { Routes } from '@angular/router';
import {ProductsListComponent} from './pages/products-list/products-list.component';
import {AuthComponent} from './pages/auth/auth.component';
import {HomeComponent} from './components/home/home.component';
import {CartComponent} from './pages/cart/cart.component';

export const routes: Routes = [{
  path: '',
  pathMatch: 'full',
  component: HomeComponent,
},
  {
    path: 'auth',
    component: AuthComponent,
  },
  {
    path: 'auth',
    component: AuthComponent,
  },
  {
    path: 'product',
    component: ProductsListComponent,
  },
  {
    path: 'cart',
    component: CartComponent,
  }
];
