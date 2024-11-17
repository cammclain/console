import { useState } from 'react';
import logo from './assets/images/logo-universal.png';
import './App.css';
import { Greet } from "../wailsjs/go/main/App";
import { useAuth } from './hooks/useAuth';
import { Button } from './components/ui/button';
import { Bell, Settings, User } from 'lucide-react';
import { Sidebar } from './components/ui/sidebar';
import { useIsMobile } from './hooks/use-mobile';
import { DashboardLayout } from './components/layout/Dashboard';

function App() {
    const { isAuthenticated, setIsAuthenticated } = useAuth();
    const isMobile = useIsMobile();
    const [activeTab, setActiveTab] = useState('dashboard');


    if (!isAuthenticated) {
        return (
            <div className="flex min-h-screen items-center justify-center">
                <div className="w-full max-w-md space-y-8 p-8">
                    <div className="text-center">
                        <img src={logo} className="mx-auto h-12 w-auto" alt="logo" />
                        <h2 className="mt-6 text-3xl font-bold">Sign in to your account</h2>
                    </div>
                    <form className="mt-8 space-y-6">
                        <div className="space-y-4">
                            <input type="text" placeholder="Server URL" className="w-full rounded-md border p-2" />
                            <input type="text" placeholder="Username" className="w-full rounded-md border p-2" />
                            <input type="password" placeholder="Password" className="w-full rounded-md border p-2" />
                            <Button 
                                className="w-full" 
                                onClick={() => setIsAuthenticated(true)}
                            >
                                Sign in
                            </Button>
                        </div>
                    </form>
                </div>
            </div>
        );
    }

    return (
        <div className="flex h-screen">
            <DashboardLayout />
            <div className="flex-1 p-8">
                <div className="flex justify-end space-x-4 mb-8">
                    <Button variant="ghost" size="icon">
                        <Bell className="h-5 w-5" />
                    </Button>
                    <Button variant="ghost" size="icon">
                        <User className="h-5 w-5" />
                    </Button>
                    <Button variant="ghost" size="icon">
                        <Settings className="h-5 w-5" />
                    </Button>
                </div>

                <main className="h-full">
                    {/* Content for each tab will go here */}
                    <div className="text-2xl font-bold">
                        {activeTab.charAt(0).toUpperCase() + activeTab.slice(1)}
                    </div>
                </main>
            </div>
        </div>
    );
}

export default App
