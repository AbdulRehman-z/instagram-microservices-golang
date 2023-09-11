import { app } from "./app";
import { prisma } from "./services/prisma.service";

app.listen(3000, () => {
  prisma.$connect();
  console.log("Auth-service is running on port 3000");
});
