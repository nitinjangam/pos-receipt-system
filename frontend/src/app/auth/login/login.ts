import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { AuthService } from '../../api/api/auth.service';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-login',
  standalone: true,
  imports: [FormsModule, CommonModule],
  templateUrl: './login.html',
  styleUrls: ['./login.scss']
})
export class LoginComponent {
  username: string = '';
  password: string = '';

  constructor(private authService: AuthService) {}

  onLogin() {
    console.log('Login button clicked:', this.username, this.password);
    
    this.authService.authLoginPost({ username: this.username, password: this.password })
      .subscribe({
        next: (res) => console.log('Login success:', res),
        error: (err) => console.error('Login failed:', err)
      });
  }
}
