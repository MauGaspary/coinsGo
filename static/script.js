// Estado da aplicação
let auth = {
    username: null,
    token: null
};

// Event Listeners
document.getElementById('loginForm').addEventListener('submit', handleLogin);

// Funções
async function handleLogin(e) {
    e.preventDefault();
    
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const errorEl = document.getElementById('loginError');
    const submitBtn = e.target.querySelector('button[type="submit"]');
    
    // Limpar erro anterior e redefinir a cor
    errorEl.textContent = '';
    errorEl.style.color = '#ff4c4c'; // cor de erro padrão
    
    try {
        submitBtn.disabled = true;
        submitBtn.textContent = 'Autenticando...';

        // 1. Fazer o post para /login
        const loginResponse = await fetch('/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ account_id: username, password: password })
        });

        if (!loginResponse.ok) {
            throw new Error('Usuário ou senha inválidos!');
        }

        const data = await loginResponse.json();
        
        if (!data.token) {
            throw new Error('Token não foi retornado pelo servidor.');
        }

        // Armazenar credenciais
        auth.username = username;
        auth.token = data.token;

        // 2. Mostrar mensagem de sucesso específica
        errorEl.style.color = '#4CAF50'; // cor de sucesso (verde)
        errorEl.textContent = 'Login efetuado com sucesso! Redirecionando...';

        // 3. Aguardar um instante para o usuário ver a mensagem, e então buscar os dados
        setTimeout(async () => {
            const hasAccess = await loadBalance();
            if (hasAccess) {
                showDashboard();
            } else {
                errorEl.style.color = '#ff4c4c';
                errorEl.textContent = 'Sessão expirada ou inválida. Faça login novamente!';
                auth.username = null;
                auth.token = null;
            }
            submitBtn.disabled = false;
            submitBtn.textContent = 'Entrar';
        }, 1200);

    } catch (error) {
        errorEl.style.color = '#ff4c4c';
        errorEl.textContent = error.message;
        auth.username = null;
        auth.token = null;
        submitBtn.disabled = false;
        submitBtn.textContent = 'Entrar';
    }
}

async function loadBalance() {
    const errorEl = document.getElementById('errorMessage');
    errorEl.textContent = '';
    
    if (!auth.username || !auth.token) {
        errorEl.textContent = 'Faça login primeiro!';
        return false;
    }
    
    try {
        // Remover account_id da URL, a API pega do JWT!
        const response = await fetch(
            `/account/balance`,
            {
                headers: {
                    'Authorization': `Bearer ${auth.token}`
                }
            }
        );
        
        if (response.status === 401 || response.status === 400) {
            return false;
        }
        
        if (!response.ok) {
            throw new Error('Erro ao buscar saldo');
        }
        
        const data = await response.json();
        
        // Formatar saldo
        const balance = data.balance || 0;
        const formatted = new Intl.NumberFormat('pt-BR', {
            style: 'currency',
            currency: 'BRL'
        }).format(balance);
        
        document.getElementById('balanceAmount').textContent = formatted;
        return true;
        
    } catch (error) {
        errorEl.textContent = 'Erro: ' + error.message;
        return false;
    }
}

function showDashboard() {
    // Esconder login
    document.getElementById('loginScreen').classList.remove('active');
    
    // Mostrar dashboard
    document.getElementById('dashboardScreen').classList.add('active');
    
    // Atualizar nome (Corrigido para utilizar o ID correto welcomeUsername)
    document.getElementById('welcomeUsername').textContent = auth.username;
}

function logout() {
    // Limpar credenciais
    auth.username = null;
    auth.token = null;
    
    // Limpar form
    document.getElementById('loginForm').reset();
    document.getElementById('loginError').textContent = '';
    document.getElementById('errorMessage').textContent = '';
    
    // Mostrar login
    document.getElementById('loginScreen').classList.add('active');
    document.getElementById('dashboardScreen').classList.remove('active');
}
