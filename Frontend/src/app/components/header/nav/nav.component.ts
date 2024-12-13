import { Component } from '@angular/core';
import {PrimaryButtonComponent} from '../../primary-button/primary-button.component';
import {RouterLink} from '@angular/router';

@Component({
  selector: 'app-nav',
  standalone: true,
  imports: [
    PrimaryButtonComponent,
    RouterLink
  ],
  templateUrl: './nav.component.html',
  styles: ``
})
export class NavComponent {

}
