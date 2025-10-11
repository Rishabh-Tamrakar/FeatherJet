// FeatherJet Demo JavaScript

// API base URL
const API_BASE = '/api';

// DOM elements
const apiResponseElement = document.getElementById('api-response');

// Test the hello API endpoint
async function testAPI() {
    try {
        showLoading();
        const response = await fetch(`${API_BASE}/hello`);
        const data = await response.json();
        displayResponse(data, 'Hello API Response');
    } catch (error) {
        displayError('Failed to fetch hello API', error);
    }
}

// Get server information
async function getServerInfo() {
    try {
        showLoading();
        const response = await fetch(`${API_BASE}/info`);
        const data = await response.json();
        displayResponse(data, 'Server Info Response');
    } catch (error) {
        displayError('Failed to fetch server info', error);
    }
}

// Check server status
async function checkStatus() {
    try {
        showLoading();
        const response = await fetch(`${API_BASE}/status`);
        const data = await response.json();
        displayResponse(data, 'Server Status Response');
    } catch (error) {
        displayError('Failed to fetch server status', error);
    }
}

// Display API response in the response box
function displayResponse(data, title = 'API Response') {
    const formattedData = JSON.stringify(data, null, 2);
    apiResponseElement.textContent = `${title}:\n\n${formattedData}`;
    apiResponseElement.style.color = '#a0aec0';
}

// Display error message
function displayError(message, error) {
    const errorText = `Error: ${message}\n\nDetails: ${error.message || error}`;
    apiResponseElement.textContent = errorText;
    apiResponseElement.style.color = '#fc8181';
}

// Show loading state
function showLoading() {
    apiResponseElement.textContent = 'Loading...';
    apiResponseElement.style.color = '#a0aec0';
}

// Theme management
function setTheme(theme) {
    document.documentElement.setAttribute('data-theme', theme);
    localStorage.setItem('theme', theme);
    
    const icon = document.getElementById('theme-toggle-icon');
    const text = document.getElementById('theme-toggle-text');
    
    if (theme === 'dark') {
        icon.textContent = '🌙';
        text.textContent = 'Dark';
    } else {
        icon.textContent = '🌞';
        text.textContent = 'Light';
    }
}

function toggleTheme() {
    const currentTheme = localStorage.getItem('theme') || 'light';
    const newTheme = currentTheme === 'light' ? 'dark' : 'light';
    setTheme(newTheme);
}

// Add some interactivity to the page
document.addEventListener('DOMContentLoaded', function() {
    // Initialize theme
    const savedTheme = localStorage.getItem('theme') || 'light';
    setTheme(savedTheme);
    
    // Add theme toggle event listener
    const themeToggle = document.getElementById('theme-toggle');
    themeToggle.addEventListener('click', toggleTheme);
    // Smooth scrolling for navigation links
    const navLinks = document.querySelectorAll('.nav-links a[href^="#"]');
    navLinks.forEach(link => {
        link.addEventListener('click', function(e) {
            e.preventDefault();
            const targetId = this.getAttribute('href').substring(1);
            const targetElement = document.getElementById(targetId);
            if (targetElement) {
                targetElement.scrollIntoView({
                    behavior: 'smooth'
                });
            }
        });
    });

    // Add a welcome message to the console
    console.log('🚀 Welcome to FeatherJet Demo!');
    console.log('This is a lightweight web server built with Go.');
    console.log('Try the API endpoints by clicking the buttons above.');

    // Add keyboard shortcuts
    document.addEventListener('keydown', function(e) {
        // Ctrl/Cmd + 1: Test API
        if ((e.ctrlKey || e.metaKey) && e.key === '1') {
            e.preventDefault();
            testAPI();
        }
        // Ctrl/Cmd + 2: Server Info
        if ((e.ctrlKey || e.metaKey) && e.key === '2') {
            e.preventDefault();
            getServerInfo();
        }
        // Ctrl/Cmd + 3: Server Status
        if ((e.ctrlKey || e.metaKey) && e.key === '3') {
            e.preventDefault();
            checkStatus();
        }
    });

    // Show keyboard shortcuts hint
    setTimeout(() => {
        console.log('💡 Keyboard shortcuts:');
        console.log('  Ctrl/Cmd + 1: Test Hello API');
        console.log('  Ctrl/Cmd + 2: Get Server Info');
        console.log('  Ctrl/Cmd + 3: Check Server Status');
    }, 1000);
});

// Feature card animations
function animateFeatureCards() {
    const cards = document.querySelectorAll('.feature-card');
    const observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.style.opacity = '1';
                entry.target.style.transform = 'translateY(0)';
            }
        });
    });

    cards.forEach(card => {
        card.style.opacity = '0';
        card.style.transform = 'translateY(20px)';
        card.style.transition = 'opacity 0.6s ease, transform 0.6s ease';
        observer.observe(card);
    });
}

// Initialize animations when page loads
document.addEventListener('DOMContentLoaded', animateFeatureCards);