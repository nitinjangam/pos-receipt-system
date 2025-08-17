import { Component, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { LoginComponent } from './auth/login/login';

@Component({
  selector: 'app-root',
  imports: [LoginComponent],
  template: '<app-login></app-login>',
  styleUrl: './app.scss'
})
export class App {
  protected readonly title = signal('frontend');
}
