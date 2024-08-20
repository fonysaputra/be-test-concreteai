import cron from 'node-cron';
import { PrismaClient } from '@prisma/client';
const prisma = new PrismaClient();

// Function to start the cron jobs
export const startCronJobs = (): void => {
  cron.schedule('*/30 * * * * *', async () => {
    console.log('Running a task every 30 seconds');
  
    try {
      // Query transactions with a specific status
      const statusToFilter = 'pending'; 
      const newStatus = 'success'
      const transactions  = await prisma.transactions.findMany({
        where: {
          status: statusToFilter,
        },
      });
      console.log('Fetched Transactions with status', statusToFilter, ':', transactions);

      // Update the status of the fetched transactions
      if (transactions.length > 0) {
        for (const transaction of transactions) {
          const updatedTransaction = await prisma.transactions.updateMany({
            where: {
              status: statusToFilter,
              id: transaction.id,
            },
            data: {
              status: newStatus,
            },
          });
          console.log('Updated Transaction:', updatedTransaction);
        }
      } else {
        console.log('No transactions found with status', statusToFilter);
      }

    } catch (error) {
      console.error('Error fetching transactions:', error);
    }
  });
};