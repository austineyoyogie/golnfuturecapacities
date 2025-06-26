import {AfterViewInit, Component} from '@angular/core';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements AfterViewInit {
  ngAfterViewInit(): void {
    (() => {
      const togglePassword = document.querySelector('#togglePassword');
      const password = document.querySelector('#password');
      togglePassword.addEventListener('click', () => {
        const  type = password.getAttribute('type') === 'password' ? 'text' : 'password';
        password.setAttribute('type', type);
        password.classList.toggle('bi-eye');

      });
    })();
  }
}
