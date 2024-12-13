import {Component, output} from '@angular/core';
import {NgOptimizedImage} from '@angular/common';
import {PrimaryButtonComponent} from '../primary-button/primary-button.component';
import {Router} from '@angular/router';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [],
  templateUrl: './home.component.html',
  styles: ``
})
export class HomeComponent {
  btnClicked = output()
}
