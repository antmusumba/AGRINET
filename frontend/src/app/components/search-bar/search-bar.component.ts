import {Component, EventEmitter, Output} from '@angular/core';
import {FormsModule} from '@angular/forms';

@Component({
  selector: 'app-search-bar',
  standalone: true,
  imports: [
    FormsModule
  ],
  templateUrl: './search-bar.component.html',
  styles: ``
})
export class SearchBarComponent {
  searchQuery: string = '';

  @Output() search = new EventEmitter<string>();

  onSearch() {
    this.search.emit(this.searchQuery); // Emit the search query to parent
  }
}
