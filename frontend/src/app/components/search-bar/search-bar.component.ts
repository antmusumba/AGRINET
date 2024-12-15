import { Component, EventEmitter, Output } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { RouterLink } from '@angular/router';
import nlp from 'compromise';

declare var webkitSpeechRecognition: any;

@Component({
  selector: 'app-search-bar',
  standalone: true,
  imports: [FormsModule, RouterLink, ReactiveFormsModule],
  templateUrl: './search-bar.component.html',
})
export class SearchBarComponent {
  searchQuery: string = '';
  @Output() search = new EventEmitter<string>();

  recognition: any = null;

  constructor() {
    if ('webkitSpeechRecognition' in window) {
      this.recognition = new webkitSpeechRecognition();
      this.recognition.continuous = false;
      this.recognition.lang = 'en-US';
      this.recognition.interimResults = false;
      this.recognition.maxAlternatives = 1;
    }
  }

  // Start speech recognition when input is focused
  startSpeechRecognition() {
    if (this.recognition) {
      this.recognition.start();

      this.recognition.onresult = (event: any) => {
        const speechResult = event.results[0][0].transcript;
        this.searchQuery = speechResult;
        this.onSearch();
      };

      this.recognition.onerror = (event: any) => {
        console.error('Speech recognition error:', event.error);
      };
    }
  }

  // Trigger search when user types or when speech result is received
  onSearch() {
    const processedQuery = this.processQueryWithNLP(this.searchQuery);
    this.search.emit(processedQuery);
  }

  // NLP processing: Extract nouns from query
  processQueryWithNLP(query: string): string {
    const doc = nlp(query);
    const nouns = doc.nouns().out('array');
    return nouns.join(' ') || query;
  }
}
