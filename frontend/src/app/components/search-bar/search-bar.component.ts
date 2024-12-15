import { Component, EventEmitter, Output } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { RouterLink } from '@angular/router';
import nlp from 'compromise';

@Component({
  selector: 'app-search-bar',
  standalone: true,
  imports: [FormsModule, RouterLink],
  templateUrl: './search-bar.component.html',
})
export class SearchBarComponent {
  searchQuery: string = '';

  @Output() search = new EventEmitter<string>();

  // Process query using NLP before emitting
  onSearch() {
    const processedQuery = this.processQueryWithNLP(this.searchQuery);
    this.search.emit(processedQuery);
  }

  // Example NLP processing
  processQueryWithNLP(query: string): string {
    const doc = nlp(query);

    const nouns = doc.nouns().out('array');

    return nouns.join(' ') || query;
  }
}
