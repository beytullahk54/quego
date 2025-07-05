"use client"

import { Dialog, DialogContent, DialogTitle, DialogTrigger } from "@/components/ui/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { useState } from "react";
import { AuthService } from "@/lib/auth/AuthService";
import { useUserStore } from "@/store/useUserStore";
import { toast } from "react-toastify";

interface UpdateCreateFormProps {
  onClose: () => void;
  open: boolean;
  job: any;
}

interface JobFormData {
  url: string;
  method: string;
  execute_at: string;
}

const UpdateCreateForm = ({ onClose, open, job }: UpdateCreateFormProps) => {

    const { jobs, isLoading, error, fetchUsers } = useUserStore();

    const auth = AuthService.getInstance();
    const formatDateForInput = (dateString: string) => {
        if (!dateString) return '';
        const date = new Date(dateString);
        const localDate = new Date(date.getTime() - (date.getTimezoneOffset() * 60000));
        return localDate.toISOString().slice(0, 16);
    };

    const [formData, setFormData] = useState<JobFormData>({
        url: job.url,
        method: job.method,
        execute_at: formatDateForInput(job.execute_at)
    });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData(prev => ({
            ...prev,
            [name]: value
        }));
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        console.log('Form submitted:', formData);
        
        const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/jobs/${job.id}`, {
            method: 'PUT',
            headers: {
              'Accept': 'application/json',
              'Authorization': `Bearer ${auth.getToken()}`
            },
            body: JSON.stringify(formData),
        }).then((response) => {
            if (response.ok) {
                fetchUsers(auth.getToken() || '');
                toast.success('İş başarıyla güncellendi!');
            }
        });
        
        onClose();
    };

    return (
        <Dialog open={open} onOpenChange={onClose}>
            <DialogTrigger asChild>
                <Button variant="outline">Güncelle</Button>
            </DialogTrigger>
            <DialogContent>
                <DialogTitle>Güncelle</DialogTitle>
                <form onSubmit={handleSubmit} className="space-y-4">
                    <div className="space-y-2">
                        <Label htmlFor="url">URL</Label>
                        <Input
                            id="url"
                            name="url"
                            type="url"
                            value={formData.url}
                            onChange={handleChange}
                            placeholder="https://example.com/api/endpoint"
                            required
                        />
                    </div>
                    
                    <div className="space-y-2">
                        <Label htmlFor="method">HTTP Method</Label>
                        <Select 
                            value={formData.method}
                            onValueChange={(value) => setFormData({...formData, method: value})}
                        >
                            <SelectTrigger>
                                <SelectValue placeholder="Method seçin" />
                            </SelectTrigger>
                            <SelectContent>
                                <SelectItem value="GET">GET</SelectItem>
                                <SelectItem value="POST">POST</SelectItem>
                                <SelectItem value="PUT">PUT</SelectItem>
                                <SelectItem value="DELETE">DELETE</SelectItem>
                                <SelectItem value="PATCH">PATCH</SelectItem>
                            </SelectContent>
                        </Select>
                    </div>
                    
                    <div className="space-y-2">
                        <Label htmlFor="execute_at">Çalıştırma Zamanı</Label>
                        <Input
                            id="execute_at"
                            name="execute_at"
                            type="datetime-local"
                            value={formData.execute_at}
                            onChange={handleChange}
                            required
                        />
                    </div>
                    
                    <div className="flex justify-end space-x-2 pt-4">
                        <Button type="button" variant="outline" onClick={onClose}>
                            İptal
                        </Button>
                        <Button type="submit">
                            Güncelle
                        </Button>   
                    </div>
                </form>
            </DialogContent>
        </Dialog>
    );
};

export default UpdateCreateForm;
