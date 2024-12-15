import { ComponentFixture, TestBed } from '@angular/core/testing';
import { SearchBarComponent } from './search-bar.component';
import { FormsModule } from '@angular/forms';
import * as nlp from 'compromise';

describe('SearchBarComponent', () => {
  let component: SearchBarComponent;
  let fixture: ComponentFixture<SearchBarComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SearchBarComponent, FormsModule],
    }).compileComponents();

    fixture = TestBed.createComponent(SearchBarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should emit processed query when onSearch is called', () => {
    spyOn(component.search, 'emit');

    component.searchQuery = 'Buy apple and banana';
    component.onSearch();

    expect(component.search.emit).toHaveBeenCalledWith('test query');
  });

  it('should process the query with NLP and extract nouns', () => {
    const nlpMock = jasmine.createSpyObj('nlp', ['nouns']);

    nlpMock.nouns.and.returnValue({
      out: () => ['test', 'query'],
    });

    spyOn(nlp, 'default').and.returnValue(nlpMock);

    const processedQuery = component.processQueryWithNLP(
      'Buy apple and banana'
    );

    expect(processedQuery).toBe('test query');
  });

  it('should return original query if no nouns are extracted', () => {
    const nlpMock = jasmine.createSpyObj('nlp', ['nouns']);

    nlpMock.nouns.and.returnValue({
      out: () => [],
    });

    spyOn(nlp, 'default').and.returnValue(nlpMock);

    const processedQuery = component.processQueryWithNLP('something else');

    expect(processedQuery).toBe('something else');
  });
});
