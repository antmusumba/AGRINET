import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-primary-button',
  standalone: true,
  imports: [],
  templateUrl: './primary-button.component.html',
})
export class PrimaryButtonComponent {
  @Input() label: string = '';
  @Output() btnClicked = new EventEmitter<void>();

  onButtonClick() {
    this.btnClicked.emit();
  }
}
