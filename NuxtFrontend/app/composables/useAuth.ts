
// Definimos la estructura de nuestro usuario 
interface User {
    email: string;
    id?: number;
    profile_image_path?: string;
}

export const useAuth = () => {
    // useCookie es una función mágica de Nuxt.
    // Guarda el dato en una cookie del navegador automáticamente.
    // Si recargas la página, el dato sigue ahí.
    const token = useCookie<string | null>('auth_token'); // auth_token es el nombre de la cookie
    
    // useState es para manejar estado reactivo dentro de la app (como variables globales)
    const user = useState<User | null>('auth_user', () => null);

    // Función para guardar sesión (Login/Register)
    const setToken = (newToken: string) => {
        token.value = newToken;
    };

    // Función para cerrar sesión
    const logout = () => {
        token.value = null;
        user.value = null;
        // Redirigir al login
        navigateTo('/login');
    };

    return {
        token,
        user,
        setToken,
        logout
    };
};