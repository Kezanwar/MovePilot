import React, { type FC, useState, useRef, useEffect } from 'react';
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetHeader,
  SheetTitle,
  SheetFooter
} from '@app/components/ui/sheet';
import { Button } from '@app/components/ui/button';

type Props = {
  open: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  onAccept: () => void;
};

const TermsAndConditionsModal: FC<Props> = ({ open, setOpen, onAccept }) => {
  const [hasScrolledToBottom, setHasScrolledToBottom] = useState(false);
  const scrollContainerRef = useRef<HTMLDivElement>(null);

  const handleScroll = () => {
    const container = scrollContainerRef.current;
    if (!container) return;

    const isAtBottom = 
      container.scrollHeight - container.scrollTop <= container.clientHeight + 10;
    
    if (isAtBottom) {
      setHasScrolledToBottom(true);
    }
  };

  useEffect(() => {
    if (!open) {
      setHasScrolledToBottom(false);
    }
  }, [open]);

  return (
    <Sheet open={open} onOpenChange={setOpen}>
      <SheetContent side="right" className="rounded">
        <SheetHeader>
          <SheetTitle>Terms and Conditions</SheetTitle>
          <SheetDescription>
            Please read and accept our terms and conditions to continue
          </SheetDescription>
        </SheetHeader>

        <div 
          ref={scrollContainerRef}
          onScroll={handleScroll}
          className="mx-4 h-[calc(100%-140px)] overflow-y-auto rounded-lg bg-slate-100 p-4 dark:border dark:bg-inherit">
          <div className="space-y-4 text-sm">
            <h3 className="text-base font-semibold">1. Acceptance of Terms</h3>
            <p className="text-muted-foreground">
              By accessing and using this service, you accept and agree to be
              bound by the terms and provision of this agreement. If you do not
              agree to abide by the above, please do not use this service.
            </p>

            <h3 className="text-base font-semibold">2. Use License</h3>
            <p className="text-muted-foreground">
              Permission is granted to temporarily download one copy of the
              materials (information or software) on our service for personal,
              non-commercial transitory viewing only. This is the grant of a
              license, not a transfer of title.
            </p>

            <h3 className="text-base font-semibold">3. Disclaimer</h3>
            <p className="text-muted-foreground">
              The materials on our service are provided on an 'as is' basis. We
              make no warranties, expressed or implied, and hereby disclaim and
              negate all other warranties including, without limitation, implied
              warranties or conditions of merchantability, fitness for a
              particular purpose, or non-infringement of intellectual property
              or other violation of rights.
            </p>

            <h3 className="text-base font-semibold">4. Limitations</h3>
            <p className="text-muted-foreground">
              In no event shall our company or its suppliers be liable for any
              damages (including, without limitation, damages for loss of data
              or profit, or due to business interruption) arising out of the use
              or inability to use the materials on our service, even if we or
              our authorized representative has been notified orally or in
              writing of the possibility of such damage.
            </p>

            <h3 className="text-base font-semibold">5. Privacy Policy</h3>
            <p className="text-muted-foreground">
              Your privacy is important to us. Our privacy policy explains how
              we collect, use, and protect your information when you use our
              service. By using our service, you agree to the collection and use
              of information in accordance with our privacy policy.
            </p>

            <h3 className="text-base font-semibold">6. User Accounts</h3>
            <p className="text-muted-foreground">
              When you create an account with us, you must provide information
              that is accurate, complete, and current at all times. You are
              responsible for safeguarding the password and for all activities
              that occur under your account.
            </p>

            <h3 className="text-base font-semibold">7. Prohibited Uses</h3>
            <p className="text-muted-foreground">
              You may not use our service: (a) for any unlawful purpose or to
              solicit others to perform unlawful acts; (b) to violate any
              international, federal, provincial, or state regulations, rules,
              laws, or local ordinances; (c) to infringe upon or violate our
              intellectual property rights or the intellectual property rights
              of others.
            </p>

            <h3 className="text-base font-semibold">8. Termination</h3>
            <p className="text-muted-foreground">
              We may terminate or suspend your account immediately, without
              prior notice or liability, for any reason whatsoever, including
              without limitation if you breach the Terms.
            </p>

            <h3 className="text-base font-semibold">9. Governing Law</h3>
            <p className="text-muted-foreground">
              These Terms shall be governed and construed in accordance with the
              laws of our jurisdiction, without regard to its conflict of law
              provisions. Our failure to enforce any right or provision of these
              Terms will not be considered a waiver of those rights.
            </p>

            <h3 className="text-base font-semibold">10. Changes to Terms</h3>
            <p className="text-muted-foreground">
              We reserve the right, at our sole discretion, to modify or replace
              these Terms at any time. If a revision is material, we will try to
              provide at least 30 days notice prior to any new terms taking
              effect.
            </p>

            <h3 className="text-base font-semibold">11. Contact Information</h3>
            <p className="text-muted-foreground">
              If you have any questions about these Terms, please contact us at
              support@example.com.
            </p>
          </div>
        </div>

        <SheetFooter className="flex flex-row justify-end gap-2 border-t">
          <Button variant="outline" onClick={() => setOpen(false)}>
            Cancel
          </Button>
          <Button 
            onClick={onAccept} 
            type="button"
            disabled={!hasScrolledToBottom}
          >
            Accept Terms & Continue
          </Button>
        </SheetFooter>
      </SheetContent>
    </Sheet>
  );
};

export default TermsAndConditionsModal;
