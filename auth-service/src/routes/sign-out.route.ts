import express, { Request, Response, NextFunction } from "express";
import { supabase } from "../services/supabase.service";
import { BadRequestError } from "@underthehoodjs/commonjs";

const router = express.Router();

router.post(
  "/api/users/sign-out",
  async (req: Request, res: Response, next: NextFunction) => {
    try {
      const { error } = await supabase.auth.signOut();
      if (error) {
        throw new BadRequestError(error.message);
      }
      res.status(200).json({
        message: "success",
      });
    } catch (error) {
      next(error);
    }
  }
);

export { router as signOutRouter };
