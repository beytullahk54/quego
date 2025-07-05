"use client"

import { Table, TableCaption, TableHeader, TableRow, TableHead, TableBody, TableCell } from "@/components/ui/table"
import { useUserStore } from "@/store/useUserStore";
import { useEffect, useCallback } from "react";
import { Button } from "@/components/ui/button";
import { CheckCircle2Icon, Pencil, Trash2 } from "lucide-react";
import { AuthService } from "@/lib/auth/AuthService";
import { useState } from "react";
import JobCreateForm from "./_form/JobCreateForm";
import UpdateCreateForm from "./_form/UpdateCreateForm";
import { toast } from "react-toastify";
import Swal from 'sweetalert2'
import { IoIosInformationCircleOutline } from "react-icons/io";
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "@/components/ui/tooltip";

export default function Jobs() {
  const { jobs, isLoading, error, fetchUsers } = useUserStore();
  const auth = AuthService.getInstance();
  const [open, setOpen] = useState(false);
  const [updateOpen, setUpdateOpen] = useState(false);
  const [updateData, setUpdateData] = useState<any>(null);

  const fetchUsersCallback = useCallback(() => {
    const token = auth.getToken();
    fetchUsers(token || '');
    
  }, [ fetchUsers]);

  useEffect(() => {
    fetchUsersCallback();
  }, [fetchUsersCallback]);

  const handleDelete = (id: number) => {
    const token = auth.getToken();

    Swal.fire({
      title: 'Bu işi silmek istediğinize emin misiniz?',
      showDenyButton: false,
      showCancelButton: true,
      confirmButtonText: 'Evet',
      cancelButtonText: 'Vazgeç',
    }).then((result) => {
      if (result.isConfirmed) {
        fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/jobs/${id}`, {
          method: 'DELETE',
          headers: {
            'Accept': 'application/json',
            'Authorization': `Bearer ${token}`
          }
        }).then((response) => {
          if (response.ok) {
            fetchUsersCallback();
            toast.success('İş başarıyla silindi!');
          }
        });
      } else if (result.isDenied) {
        toast.info('İş silme işlemi iptal edildi!');
      }
    })

  
  };

  const handleUpdate = (job: any) => {
    setUpdateOpen(true);
    setUpdateData(job);
  };
  
  if (isLoading) {
    return (
      <>
        <h1 className="text-3xl font-bold mb-8">İşler</h1>
        <div>Yükleniyor...</div>
      </>
    );
  }

  if (error) {
    return (
      <>
        <h1 className="text-3xl font-bold mb-8">İşler</h1>
        <div className="text-red-500">{error}</div>
      </>
    );
  }
  
  return (
    <>
      <div className="flex justify-between items-center mb-8">
        <div className="flex justify-between ">
          <h1 className="text-3xl font-bold display-inline">
            İşler
          </h1>
          <Tooltip>
            <TooltipTrigger asChild>
              <IoIosInformationCircleOutline  className="inline ml-2 "/>
            </TooltipTrigger>
            <TooltipContent>
              <p>İşler belirlediğiniz tarih ve saatte çalışacaktır</p>
            </TooltipContent>
          </Tooltip>
        </div>
        <Button onClick={() => setOpen(true)}>
          Yeni İş Ekle 
        </Button>
      </div>
      <div className="rounded-md border">
        <Table>
          <TableCaption>Toplam {jobs.length} iş</TableCaption>
          
          <TableHeader>
            <TableRow>
              <TableHead>URL</TableHead>
              <TableHead>Method</TableHead>
              <TableHead>Tarih / Saat</TableHead>
              <TableHead className="text-right">İşlemler</TableHead>
            </TableRow>
          </TableHeader>
          
          <TableBody>
            {jobs.map((job) => (
              <TableRow key={job.id}>
                <TableCell className="font-medium">{job.url}</TableCell>
                <TableCell>{job.method}</TableCell>
                <TableCell>
                  {new Date(job.execute_at).toLocaleDateString('tr-TR')} / {new Date(job.execute_at).toLocaleTimeString('tr-TR')}
                </TableCell>
                <TableCell className="text-right">
                  <div className="flex justify-end gap-2">
                    <Button variant="outline" size="icon" onClick={() => handleUpdate(job)}>
                      <Pencil className="h-4 w-4" />
                    </Button>
                    <Button variant="outline" size="icon" className="text-red-500 hover:text-red-700" onClick={() => handleDelete(job.id)} >
                      <Trash2 className="h-4 w-4" />
                    </Button>
                  </div>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
      {open && <JobCreateForm onClose={() => setOpen(false)} open={open} />}
      {updateOpen && <UpdateCreateForm onClose={() => setUpdateOpen(false)} open={updateOpen} job={updateData} />}
    </>
  );
}
