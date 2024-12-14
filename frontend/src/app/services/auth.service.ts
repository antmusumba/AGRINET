import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BehaviorSubject, Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { environment } from '../environment/environment.dev';

export interface User {
  id: string;
  email: string;
  firstName: string;
  lastName: string;
  phone: string;
}

export interface AuthResponse {
  user: User;
  token: string;
}

export interface RegisterRequest {
  firstName: string;
  lastName: string;
  email: string;
  phone: string;
  password: string;
}

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private readonly API_URL = `${environment.apiUrl}/api/auth`;

  private currentUserSubject = new BehaviorSubject<User | null>(null);
  currentUser$ = this.currentUserSubject.asObservable();

  private tokenSubject = new BehaviorSubject<string | null>(null);

  constructor(private http: HttpClient) {}

  login(email: string, password: string): Observable<AuthResponse> {
    return this.http
      .post<AuthResponse>(`${this.API_URL}/login`, {
        email,
        password,
      })
      .pipe(
        map((response: AuthResponse) => {
          this.tokenSubject.next(response.token);
          this.currentUserSubject.next(response.user);
          return response;
        })
      );
  }

  register(userData: RegisterRequest): Observable<void> {
    return this.http
      .post<AuthResponse>(`${this.API_URL}/register`, userData)
      .pipe(map(() => void 0));
  }

  logout(): void {
    this.tokenSubject.next(null);
    this.currentUserSubject.next(null);
  }

  isLoggedIn(): boolean {
    return !!this.tokenSubject.value;
  }

  getToken(): string | null {
    return this.tokenSubject.value;
  }
}
