import { Bell, User, Settings } from 'lucide-react';
import { Button } from '../ui/button';
import { NavigationMenu, NavigationMenuList, NavigationMenuItem } from '../ui/navigation-menu';
import { useAuth, clientLogin } from '../../hooks/useAuth';

// This is the header component
// / it will contain the header
// / if the user is logged in:
// / it will contain the notifications
// / it will contain the user menu
// / it will contain the settings menu
export const Header = () => {
    const { isAuthenticated } = useAuth();

    // If the user is authenticated, return the header
    // / it will contain the notifications
    // / it will contain the user menu
    // / it will contain the settings menu
    if (!isAuthenticated) {
        return (
            <NavigationMenu>
                <NavigationMenuList>
                    <NavigationMenuItem>
                        <Button variant="link" onClick={clientLogin}>
                            Login
                        </Button>
                    </NavigationMenuItem>
                </NavigationMenuList>
            </NavigationMenu>
        );
    }
    
    if (isAuthenticated) {
        return (
            <NavigationMenu>
                <NavigationMenuList>
                    <NavigationMenuItem>
                        <Bell className="h-5 w-5" />
                    </NavigationMenuItem>
                </NavigationMenuList>
            </NavigationMenu>
        );
    }

    // If the user is not authenticated, return null
    
    return null;
};