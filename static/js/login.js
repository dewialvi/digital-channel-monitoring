const API_BASE_URL = 'http://localhost:8080/api/v1';

document
    .getElementById('loginForm')
    .addEventListener('submit', async function (e) {

        e.preventDefault();

        const email = document.getElementById('email').value;
        const password = document.getElementById('password').value;

        const alertBox = document.getElementById('alertBox');
        const loginBtn = document.getElementById('loginBtn');

        // Sembunyikan error sebelumnya
        alertBox.classList.add('d-none');

        // Disable tombol ketika proses login
        loginBtn.disabled = true;
        loginBtn.textContent = 'Memproses...';

        try {

            const response = await fetch(`${API_BASE_URL}/login`, {
                method: 'POST',

                headers: {
                    'Content-Type': 'application/json'
                },

                body: JSON.stringify({
                    email: email,
                    password: password
                })
            });

            const result = await response.json();

            if (!response.ok) {
                throw new Error(
                    result.message || 'Login gagal'
                );
            }

            // Simpan JWT Token
            localStorage.setItem(
                'token',
                result.data.token
            );

            // Simpan informasi user
            localStorage.setItem(
                'user',
                JSON.stringify(result.data.user)
            );

            // Redirect ke dashboard
            window.location.href = '/dashboard';

        } catch (error) {

            alertBox.textContent = error.message;

            alertBox.classList.remove('d-none');

        } finally {

            loginBtn.disabled = false;
            loginBtn.textContent = 'Login';

        }

    });