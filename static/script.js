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
    const token = document.getElementById('token').value;
    const errorEl = document.getElementById('loginError');
    
    // Limpar erro anterior
    errorEl.textContent = '';
    
    // Armazenar credenciais
    auth.username = username;
    auth.token = token;
    
    // Tentar carregar saldo para validar login
    try {
        const hasAccess = await loadBalance();
        if (hasAccess) {
            showDashboard();
        } else {
            errorEl.textContent = 'Usuário ou token inválidos!';
            auth.username = null;
            auth.token = null;
        }
    } catch (error) {
        errorEl.textContent = 'Erro ao conectar: ' + error.message;
        auth.username = null;
        auth.token = null;
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
        const response = await fetch(
            `/account/balance?account_id=${auth.username}`,
            {
                headers: {
                    'Authorization': auth.token
                }
            }
        );
        
        if (response.status === 401 || response.status === 400) {
            errorEl.textContent = 'Credenciais inválidas!';
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
    
    // Atualizar nome
    document.getElementById('username').textContent = auth.username;
    
    // Carregar saldo
    loadBalance();
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
