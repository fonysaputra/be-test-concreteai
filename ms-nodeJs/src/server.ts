//import Fastify, { FastifyInstance, FastifyReply, FastifyRequest } from 'fastify';
import { startCronJobs } from './cronJobTransaction';

// Initialize Fastify
//const fastify: FastifyInstance = Fastify({ logger: true });

// Route to get all payments
// fastify.get('/payments', async (request: FastifyRequest, reply: FastifyReply) => {
//   try {
//     const payments = await prisma.recurringPayment.findMany();
//     reply.send(payments);
//   } catch (error) {
//     fastify.log.error(error);
//     reply.status(500).send('Internal Server Error');
//   }
// });

// // Route to create a new payment
// fastify.post('/payments', async (request: FastifyRequest, reply: FastifyReply) => {
//   const { accountId, amount, interval, nextPayment } = request.body as {
//     accountId: number;
//     amount: number;
//     interval: string;
//     nextPayment: Date;
//   };

//   try {
//     const payment = await prisma.recurringPayment.create({
//       data: {
//         accountId,
//         amount,
//         interval,
//         nextPayment,
//       },
//     });
//     reply.status(201).send(payment);
//   } catch (error) {
//     fastify.log.error(error);
//     reply.status(500).send('Internal Server Error');
//   }
// });


// Start the server
const start = async (): Promise<void> => {
  try {
    // await fastify.listen({ port: 3001 });
    // fastify.log.info(`Server listening on ${fastify.server.address()}`);
    startCronJobs();
  } catch (err) {
   // fastify.log.error(err);
    process.exit(1);
  }
};

start();
